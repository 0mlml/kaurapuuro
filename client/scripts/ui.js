/**
 * @fileoverview This file contains functions that are used by the UI and other UI-related functions.
 * Notably also contains mouse and keyboard event handlers.
 */

/** @type {number} The current mouse X position */
let globalMouseX = 0;
/** @type {number} The current mouse Y position */
let globalMouseY = 0;
/** @type {Map.<number, boolean>} The current state of the mouse */
const globalMouseState = new Map();

/** Mouse move handler, updates the global mouse position */
document.addEventListener('mousemove', (event) => {
  globalMouseX = event.clientX;
  globalMouseY = event.clientY;
});

/** Mouse down handler, updates the global mouse state */
document.addEventListener('mousedown', (event) => {
  globalMouseState.set(event.button, true);
});

/** Mouse up handler, updates the global mouse state */
document.addEventListener('mouseup', (event) => {
  globalMouseState.set(event.button, false);
});

/** @type {Map<string, boolean>} The current state of the keyboard */
const globalKeyboardState = new Map();
/** @type {Map<string, Function>} Keybindings; key format: "key1+key2+key3" */
const globalKeybindings = new Map();

/**
 * @description Registers a keybinding
 * @param {string} keyCombo The key combination to bind to
 * @param {Function} func The function to call when the key combination is pressed
 */
const registerKeybinding = (keyCombo, func) => {
  // Normalize the key combo
  keyCombo = keyCombo.split('+').sort().join('+').toLowerCase();

  globalKeybindings.set(keyCombo, func);
};

/** Key down handler, updates the global keyboard state */
document.addEventListener('keydown', (event) => {
  globalKeyboardState.set(event.key.toLowerCase(), true);

  const keyCombo = [...globalKeyboardState.keys()].sort().join('+');
  const func = globalKeybindings.get(keyCombo);
  if (func) func();
});

/** Key up handler, updates the global keyboard state */
document.addEventListener('keyup', (event) => {
  globalKeyboardState.delete(event.key.toLowerCase());
});

/**
 * @description Shorthand function for creating an element and appending it to a parent
 * @param {string} type The type of element to create 
 * @param {string} className The class name(s) to add to the element 
 * @param {string} alt The alt text for the element
 * @param {HTMLElement} parent The parent element to append the new element to 
 * @returns {HTMLElement} The newly created element
 */
const createElem = (type, className = '', title, parent) => {
  const elem = document.createElement(type);
  elem.className = className;
  elem.title = title;
  parent?.appendChild(elem);
  return elem;
};

/**
 * @description Generates a modal window
 * @param {{title: string, content: string | HTMLElement, width: number, height: number, resizable: boolean}} options The options for the modal window
 * @returns {HTMLElement} The modal window element
 */
const generateModalWindow = ({
  title = 'Modal Window',
  content = 'Placeholder',
  width = 400,
  height = 300,
  resizable = false,
} = {}) => {
  if (content instanceof HTMLElement) content = content.outerHTML;

  const modal = createElem('div', 'modal-window', '', document.body);
  const header = createElem('div', 'modal-window-header', '', modal);
  const titleElem = createElem('div', 'modal-window-header-title', '', header);
  const buttons = createElem('div', 'modal-window-header-buttons', '', header);

  let x1 = 0, y1 = 0, x2 = 0, y2 = 0, dragging = false;
  header.addEventListener('mousedown', e => {
    [x2, y2, dragging] = [e.clientX, e.clientY, true];
  });
  document.addEventListener('mouseup', e => {
    dragging = e.button !== 0 ? dragging : false;
  });
  document.addEventListener('mousemove', e => {
    if (dragging) {
      [x1, y1] = [x2 - e.clientX, y2 - e.clientY];
      [x2, y2] = [e.clientX, e.clientY];
      modal.style.top = `${modal.offsetTop - y1}px`;
      modal.style.left = `${modal.offsetLeft - x1}px`;
    }
  });

  titleElem.innerText = title;
  createElem('div', 'modal-window-header-button ghost-button', 'Toggle window shimmer', buttons).addEventListener('click', () => modal.classList.toggle('ghosted'));
  createElem('div', 'modal-window-header-button close-button', 'Close window', buttons).addEventListener('click', () => modal.remove());
  createElem('div', 'modal-window-content', '', modal).innerHTML = content;

  Object.assign(modal.style, { width: `${width}px`, height: `${height}px` });
  if (resizable) modal.classList.add('resizable');

  return modal;
};