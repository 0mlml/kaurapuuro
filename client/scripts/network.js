/**
 * @fileoverview This file contains configuration, state, and functions related to networking.
 * This file only contains code for creating, maintaining, monitoring, and destroying connections.
 * Game code should be in engine.js.
 */

const state = {
  socket: null,
  serverinfo: {
  },
  autoConnectInterval: null,
}

const Packets = {
  /* UnknownPacket: debug packet for netcode development */
  0: {
    deserialize: _ => null,
    serialize: _ => null,
  },
  /* UserCommandPacket: sent by client to server to update user commands */
  1: {
    deserialize: _ => null,
    serialize: (data) => {
      const buffer = new ArrayBuffer(5);
      const view = new DataView(buffer);
      view.setUint16(0, data.entityID);
      view.setUint8(2, data.commands);
      view.setUint16(3, data.barrelYaw);
      return buffer;
    }
  },
  /* UserCommandsPacket: sent by server to client to update player keys */
  2: {
    deserialize: (data) => {
      const view = new DataView(data);
      let offset = 0;
      const numCommands = view.getUint8(0);
      offset += 1;

      const commands = [];
      for (let i = 0; i < numCommands; i++) {
        const entityID = view.getUint16(offset);
        offset += 2;
        const commandField = view.getUint8(offset);
        offset += 1;
        const barrelYaw = view.getUint16(offset);
        offset += 2;
        commands.push({
          entityID: entityID,
          commands: {
            forward: commandField & 1 != 0,
            backward: commandField & 2 != 0,
            left: commandField & 4 != 0,
            right: commandField & 8 != 0,
            fire: commandField & 16 != 0,
            altfire: commandField & 32 != 0,
          },
          barrelYaw: (barrelYaw / 10) * (Math.PI / 180),
        });
      }
      console.debug('Received UserCommandsPacket', commands)
      return commands;
    },
    serialize: _ => null,
  },
  /* MapDataPacket: sent by server to client to update map data */
  3: {
    deserialize: (data) => {
      const view = new DataView(data);
      const mapData = JSON.stringify(view);
      console.debug('Received MapDataPacket', mapData);
      return mapData;
    },
    serialize: _ => null,
  },
  /* GenerateNewLobbyPacket: sent by client to server to request a new lobby */
  4: {
    deserialize: _ => null,
    serialize: (data) => {
      const name = data.name;
      const buffer = new ArrayBuffer(name.length);
      const view = new DataView(buffer);
      for (let i = 0; i < name.length; i++) {
        view.setUint8(i, name.charCodeAt(i));
      }
      return buffer;
    },
  },
  /* GenerateNewLobbyResultPacket: sent by server to client to inform of the result */
  5: {
    deserialize: (data) => {
      const view = new DataView(data);
      let uuid = "";
      for (let i = 0; i < view.byteLength; i++) {
        uuid += String.fromCharCode(view.getUint8(i));
      }
      console.debug('Received GenerateNewLobbyResultPacket', uuid);
      return uuid;
    },
    serialize: _ => null,
  },
  /* EntityListPacket: sent by server to client to update authoritative entity list, and by proxy the state of those entities */
  6: {
    deserialize: (data) => {
      const view = new DataView(data);
      const entities = [];
      let offset = 0;
      const numTanks = view.getUint8(offset);
      offset += 1;

      for (let i = 0; i < numTanks; i++) {
        const entityID = view.getUint16(offset);
        offset += 2;
        const x = view.getFloat64(offset);
        offset += 8;
        const y = view.getFloat64(offset);
        offset += 8;
        const v_x = view.getFloat64(offset);
        offset += 8;
        const v_y = view.getFloat64(offset);
        offset += 8;
        const e_flags = view.getUint8(offset);
        offset += 1;
        const barrelYaw = view.getUint16(offset) / 10 * (Math.PI / 180);
        offset += 2;
        const t_flags = view.getUint8(offset);
        offset += 1;

        // Decode t_flags for type and team
        const types = ['common', 'fast', 'rocket', 'bounce', 'ph_five', 'ph_six', 'ph_seven', 'ph_eight'];
        const type = types[t_flags & 0x07];
        const team = (t_flags >> 4) + 1;

        entities.push({
          entityID: entityID,
          x: x,
          y: y,
          v_x: v_x,
          v_y: v_y,
          barrelYaw: barrelYaw,
          flags: {
            /* Entity flags */
            isAlive: e_flags & 1 != 0,
            isInvincible: e_flags & 2 != 0,
          },
          type: type,
        });
      }
      console.debug('Received EntityListPacket', entities);
      return entities;
    },
    serialize: _ => null,
  },
}

/**
 * @description Deserializes a packet.
 * The format is:  
 * 1 Byte - Packet type; 
 * N Bytes - Payload
 * @param {ArrayBufferLike} buffer raw packet data
 * @returns {{type: number, data: any}} Deserialized packet 
 */
const deserializePacket = (buffer) => {
  const view = new DataView(buffer);
  const packetType = view.getUint8(0);
  const packetData = buffer.slice(1);

  if (!Packets[packetType]) {
    console.error(`Received unknown packet type ${packetType}`, buffer);
  }

  const packet = {
    type: packetType,
    data: Packets[packetType].deserialize(packetData),
  }

  return packet;
}


/** @type {Map<string, Function[]>} Packet listeners */
const packetListeners = new Map();

/**
 * @description Registers a new packet listener
 * @param {string} packetName The name of the packet to listen for 
 * @param {Function} func The function to call when the packet is received 
 */
const registerNewPacketListener = (packetName, func) => {
  if (!packetListeners.has(packetName)) {
    packetListeners.set(packetName, []);
  }

  packetListeners.get(packetName).push(func);
}

/**
 * @description Deregisters a packet listener
 * @param {string} packetName The name of the packet to stop listening for
 * @param {Function} func The function to deregister
 */
const deregisterPacketListener = (packetName, func) => {
  if (!packetListeners.has(packetName)) return;

  const listeners = packetListeners.get(packetName);
  const index = listeners.indexOf(func);
  if (index === -1) return;

  listeners.splice(index, 1);
}

/**
 * @description Sends a packet to the server, if a connection exists
 * @param {{type: number, data: any}} packet
 */
const sendPacket = (packet) => {
  if (config.offlineMode.value) return;

  if (!state.socket) {
    console.warn('Attempted to send packet without a connection');
    return;
  }

  const serializer = Packets[packet.type]?.serialize;

  if (!serializer) {
    console.error(`Attempted to send unknown packet type ${packet.type}`, packet);
    return;
  }

  const payload = serializer(packet.data);
  const packetBuffer = new Uint8Array(1 + payload.byteLength);
  packetBuffer[0] = packet.type;

  packetBuffer.set(new Uint8Array(payload), 1);

  state.socket.send(packetBuffer);
}


/**
 * @description Try to open a new websocket connection to the server
 * @param {string} url The URL to connect to
 * @param {boolean} force Whether to force a new connection even if one already exists
 * @returns {Promise} A promise that resolves when the connection is open
 */
const openConnection = (url, force = false) => {
  return new Promise((resolve, reject) => {
    if (config.offlineMode.value) reject('Offline mode is enabled');

    if (state.socket && !force) {
      reject('Socket already exists');
      return;
    }

    if (force) {
      state.socket?.close();
    }

    state.socket = new WebSocket(url);
    state.socket.binaryType = 'arraybuffer';

    state.socket.addEventListener('open', () => {
      resolve();
    });

    state.socket.addEventListener('error', (err) => {
      reject(err);
    });

    state.socket.addEventListener('message', (event) => {
      const packet = deserializePacket(event.data);

      if (!packetListeners.has(packet.type)) return;

      const listeners = packetListeners.get(packet.type);

      for (const listener of listeners) {
        listener(packet.data);
      }
    });

    state.socket.addEventListener('close', () => {
      console.debug('Connection closed');
      state.socket = null;
    });
  });
}