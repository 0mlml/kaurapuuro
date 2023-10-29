/**
 * @fileoverview This file contains the map editor. 
 */

const EditorMap = {
  lines: [],
  boxes: [],
  spawnpoints: [],
}

let lineCount = 0;
const inProgressLines = [];

const openNewLineModal = () => {
  let modalTitle = "Add New Line";
  const line = { x1: "", y1: "", x2: "", y2: "", color: "", index: lineCount++ };

  if (inProgressLines.length > 0) {
    modalTitle += ` (${inProgressLines.length + 1})`;
  }

  inProgressLines.push(line);

  const modal = generateModalWindow({
    title: modalTitle,
    content: generateLinePropertiesContent(line),
    width: 500,
    height: 400,
    resizable: true,
  });

  line.modal = modal;
}

/**
 * @description Generate the HTML for the line properties modal
 * @param {{x1: number, y1: number, x2: number, y2: number, color: string, index: number}} line The line to generate the HTML for
 * @returns {string} The HTML for the line properties modal
 */
const generateLinePropertiesContent = (line) => {
  return `
    <div id="line-properties">
      X1: <input type="number" id="x1" value="${line.x1}" oninput="updateAndRenderLine(this, 'x1', ${line.index});"><br>
      Y1: <input type="number" id="y1" value="${line.y1}" oninput="updateAndRenderLine(this, 'y1', ${line.index});"><br>
      X2: <input type="number" id="x2" value="${line.x2}" oninput="updateAndRenderLine(this, 'x2', ${line.index});"><br>
      Y2: <input type="number" id="y2" value="${line.y2}" oninput="updateAndRenderLine(this, 'y2', ${line.index});"><br>
      Color: <input type="text" id="color" value="${line.color}" oninput="updateAndRenderLine(this, 'color', ${line.index});"><br>
      <button onclick="saveLine(${line.index})">Save Line</button>
    </div>
  `;
}

const updateAndRenderLine = (input, property, index) => {
  const line = inProgressLines.find(line => line.index === index);
  if (input.type === 'number') line[property] = parseInt(input.value);
  else line[property] = input.value;
  renderPreviewLine(line);
}

const saveLine = (lineIndex) => {
  const line = inProgressLines.find(line => line.index === lineIndex);
  EditorMap.lines.push(line);
  const index = inProgressLines.indexOf(line);
  if (index > -1) {
    inProgressLines.splice(index, 1);
  }
  line.modal.remove();
  updateMap(EditorMap);
}

const openGlobalObjectList = () => {
  generateModalWindow({
    title: 'All Objects',
    content: generateGlobalObjectListContent(),
    width: 700,
    height: 500,
    resizable: true,
  });
}

const generateGlobalObjectListContent = () => {
  let content = '<div id="global-object-list">';

  content += "<h3>Lines</h3>";
  content += "<ul>";
  EditorMap.lines.forEach(line => {
    content += `<li onclick="openEditLineModal(line)">${line.x1},${line.y1} to ${line.x2},${line.y2}</li>`;
  });
  content += "</ul>";

  content += "<h3>Boxes</h3>";
  content += "<ul>";
  EditorMap.boxes.forEach(box => {
    content += `<li onclick="openEditBoxModal(box)">${box.x},${box.y} - ${box.width}x${box.height}</li>`;
  });
  content += "</ul>";

  content += "<h3>SpawnPoints</h3>";
  content += "<ul>";
  EditorMap.spawnpoints.forEach(spawnpoint => {
    content += `<li onclick="openEditSpawnPointModal(spawnpoint)">${spawnpoint.x},${spawnpoint.y} - ${spawnpoint.type} - ${spawnpoint.team}</li>`;
  });
  content += "</ul>";

  content += '</div>';
  return content;
}

const openEditLineModal = (line) => {
  const modal = generateModalWindow({
    title: `Edit Line - ${line.x1},${line.y1} to ${line.x2},${line.y2}`,
    content: generateLinePropertiesContent(line),
    width: 500,
    height: 400,
    resizable: true,
  });

  line.modal = modal;
}

const openEditBoxModal = (box) => {
  const modal = generateModalWindow({
    title: `Edit Box - ${box.x},${box.y} - ${box.width}x${box.height}`,
    content: generateBoxPropertiesContent(box),
    width: 500,
    height: 400,
    resizable: true,
  });

  box.modal = modal;
}

const openEditSpawnPointModal = (spawnpoint) => {
  const modal = generateModalWindow({
    title: `Edit SpawnPoint - ${spawnpoint.x},${spawnpoint.y}`,
    content: generateSpawnPointPropertiesContent(spawnpoint),
    width: 500,
    height: 400,
    resizable: true,
  });

  spawnpoint.modal = modal;
}

const generateEditorMainMenuContent = () => {
  return `
    <div id="editor-main-menu">
      <button onclick="openNewLineModal()">Add New Line</button>
      <button onclick="openNewBoxModal()">Add New Box</button>
      <button onclick="openNewSpawnPointModal()">Add New SpawnPoint</button>
      <button onclick="openGlobalObjectList()">View & Edit All Objects</button>
    </div>
  `;
}

let mapEditorMainMenu = null;
registerKeybinding('alt+e', e => {
  e.preventDefault();

  if (!mapEditorMainMenu) {
    mapEditorMainMenu = generateModalWindow({
      title: 'Map Editor Main Menu',
      content: generateEditorMainMenuContent(),
      width: 500,
      height: 300,
      resizable: true,
    });
  } else {
    mapEditorMainMenu.remove();
    mapEditorMainMenu = null;
  }
});