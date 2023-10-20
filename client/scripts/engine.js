/** @type {HTMLCanvasElement} */
const foregroundCanvas = document.getElementById('fg-canvas');
const foregroundCtx = foregroundCanvas.getContext('2d');

foregroundCanvas.width = window.innerWidth;
foregroundCanvas.height = window.innerHeight;

/** @type {HTMLCanvasElement} */
const backgroundCanvas = document.getElementById('bg-canvas');
const backgroundCtx = backgroundCanvas.getContext('2d');

backgroundCanvas.width = window.innerWidth;
backgroundCanvas.height = window.innerHeight;

registerNewPacketListener(81, (data) => { // DrawNewPointPacket
  const { x, y, color } = data;
  foregroundCtx.fillStyle = `rgb(${color}, ${color}, ${color})`;

  foregroundCtx.beginPath();
  foregroundCtx.ellipse(x, y, 5, 5, 0, 0, 2 * Math.PI);
  foregroundCtx.fill();
});

registerKeybinding('p', () => {
  const packet = {
    type: 81,
    data: {
      x: globalMouseX,
      y: globalMouseY,
      color: Math.floor(Math.random() * 0xFF),
    }
  };

  sendPacket(packet);
});