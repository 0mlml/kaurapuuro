/**
 * @fileoverview This file contains configuration, state, and functions related to rendering.
 * All canvas rendering code should be in this file.
 */

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

/**
 * @description Function to render a line to the foreground canvas
 * @param {{x1: number, y1: number, x2: number, y2: number, color: string}} line The line to render 
 */
const renderPreviewLine = (line) => {
  foregroundCtx.beginPath();
  foregroundCtx.moveTo(line.x1, line.y1);
  foregroundCtx.lineTo(line.x2, line.y2);
  foregroundCtx.strokeStyle = line.color;
  foregroundCtx.stroke();
}

/**
 * @description Function to render a box to the foreground canvas
 * @param {{X: number, Y: number, Width: number, Height: number, Color: string}} box The box to render 
 */
const renderPreviewBox = (box) => {
  foregroundCtx.fillStyle = box.Color;
  foregroundCtx.fillRect(box.X, box.Y, box.Width, box.Height);
}

/**
 * @description Function to render a spawnpoint to the foreground canvas
 * @param {{X: number, Y: number, Type: string, Team: string}} spawnpoint The spawnpoint to render
 */
const renderPreviewSpawnpoint = (spawnpoint) => {
  spawnpoint.Team = parseInt(spawnpoint.Team);
  const MAXTEAMS = 8;

  const colors = [
    '#FF0000',
    '#00FF00',
    '#0000FF',
    '#FFFF00',
    '#FF00FF',
    '#00FFFF',
    '#FFFFFF',
    '#000000'
  ];

  const color = colors[spawnpoint.Team % MAXTEAMS];

  foregroundCtx.fillStyle = color;
  foregroundCtx.beginPath();
  foregroundCtx.ellipse(spawnpoint.X, spawnpoint.Y, 10, 10, 0, 0, 2 * Math.PI);
  foregroundCtx.fill();
}

let looping = true;
let fps = 30;
let lastFrameTime = performance.now();

const renderLoop = () => {
  const now = performance.now();
  
  requestAnimationFrame(renderLoop);
};

renderLoop();
