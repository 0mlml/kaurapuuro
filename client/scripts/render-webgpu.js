/**
 * @fileoverview This file contains alternative rendering methods for WebGPU
 * 
 */

let shouldUseWebGPU = false;
let wg_adapter = null;
let wg_device = null;
(async () => {
  if (!navigator.gpu) {
    console.warn('WebGPU not supported.')
    return;
  }

  console.info('WebGPU support detected.')

  wg_adapter = await navigator.gpu.requestAdapter();

  if (!wg_adapter) {
    console.error('Could not request WebGPU adapter.');
    return;
  }

  wg_device = await wg_adapter.requestDevice();

  if (!wg_device) {
    console.error('Could not request WebGPU device.');
    return;
  }

  shouldUseWebGPU = true;
})();