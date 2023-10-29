/**
 * @fileoverview This file contains the game engine. 
 * It is meant to match the server's engine as closely as possible.
 * Rendering should be handled in render.js.
 */

const tickrate = 30;
const delta = 30 / 1;

const Map = {
  lines: [],
  spawnpoints: [],
}

const Entities = {
  players: [],
  projectiles: [],
  explosions: [],
}

let waitingForMap = true;
registerNewPacketListener(3, (data) => {
  if (!waitingForMap) {
    console.warn('Received MapDataPacket while not waiting for map data', data);
    return;
  }

  setMap(data);
  waitingForMap = false;
});

const setMap = (map) => {
  Map.lines = map.lines;
  Map.spawnpoints = map.spawnpoints;
}

const Vector = (x, y) => {
  this.x = x;
  this.y = y;
}

const Line = (start, end) => {
  this.start = start;
  this.end = end;
}

const Box = (pos, width, height) => {
  this.pos = pos;
  this.width = width;
  this.height = height;
}

/**
 * @description Usercommand object.
 * @param {number} EID 0-255 Entity ID 
 * @param {{left: boolean, right: boolean, forward: boolean, backward: boolean, fire: boolean, altfire: boolean}} commands Move flags
 * @param {number} barrelYaw Radians 
 */
const UserCommand = (EID, commands, barrelYaw) => {
  this.EID = EID;
  this.commands = commands;
  this.barrelYaw = barrelYaw;
}

