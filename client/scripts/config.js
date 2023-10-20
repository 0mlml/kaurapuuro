/**
 * @fileoverview Contains the configuration for the client
 */

const config = {
  offlineMode: {
    key: 'offlineMode',
    type: 'checkbox',
    display: 'Offline mode',
    helptext: 'When checked, the client will make no attempts to connect to a server',
    value: true
  },
  dropdownExample: {
    key: 'dropdownExample',
    type: 'dropdown',
    display: 'Just an example dropdown',
    helptext: 'foo',
    value: 'a',
    options: ['a', 'b', 'c', 'd'],
  },
  numberExample: {
    key: 'numberExample',
    type: 'number',
    display: 'Just an example number',
    helptext: 'bar',
    value: 1000,
    min: 0,
    max: 20000,
    step: 100,
  },
  serverURL: {
    key: 'serverURL',
    type: 'text',
    display: 'Server URL',
    helptext: 'The URL of the game server to connect to',
    value: 'ws://localhost:8080/ws',
    showOnlyIf: () => !config.offlineMode.value,
    callback: () => {
      openConnection(config.serverURL.value, true).then(_ => {
        console.info('Successfully connected to server');
      }).catch(err => {
        console.error('Failed to connect to server', err);
      });
    },
  },
  nickname: {
    key: 'nickname',
    type: 'text',
    display: 'Nickname',
    helptext: 'Your nickname',
    value: 'Player',
  },
  theme: {
    key: 'theme',
    type: 'dropdown',
    display: 'Theme',
    helptext: 'The theme to use',
    value: 'light',
    options: ['dark', 'light'],
    callback: () => document.body.setAttribute('data-theme', config.theme.value),
  },
  hiddenExample: {
    key: 'hiddenExample',
    type: 'hidden',
    value: 'Sneaky way to persist data using the existing system',
  },
}

/**
 * @description Handles config changes.
 * @param {HTMLElement|{id: string, value: string}} input Event or an object with a id and value.
 */
const configChangeHandler = input => {
  const cfg = config[input.id];

  if (!cfg) return;

  if (cfg.type === 'checkbox') {
    cfg.value = input.checked;
  } else {
    cfg.value = input.value;
  }

  if (cfg.callback) cfg.callback();

  localStorage.setItem(cfg.key, cfg.value);
  configDisplayUpdater();
}

/**
 * @description Updates the display based on the showOnlyIf property.
 */
const configDisplayUpdater = () => {
  for (const key in config) {
    const cfg = config[key];
    const el = document.getElementById(cfg.key);
    if (el && cfg.showOnlyIf) {
      el.parentElement.style.display = cfg.showOnlyIf() ? 'flex' : 'none';
    }
  }
}

/**
 * @description Initializes the config with stored values.
 */
const configInitializer = () => {
  for (const key in config) {
    const stored = localStorage.getItem(config[key].key);
    if (stored) {
      if (config[key].type === 'checkbox') {
        config[key].value = stored === 'true';
      } else if (config[key].type === 'number') {
        config[key].value = Number(stored);
      } else {
        config[key].value = stored;
      }

      if (config[key].callback) config[key].callback();
    }
  }
}

configInitializer();

/**
 * @description Generate the innerHTML content based on the config.
 * @returns {string} Generated innerHTML for settings.
 */
const generateSettingsContent = () => {
  let content = '<div id="settings-content">';

  for (const key in config) {
    const cfg = config[key];
    let inputHtml = '';

    switch (cfg.type) {
      case 'checkbox':
        inputHtml = `<input type="checkbox" id="${cfg.key}" ${cfg.value ? 'checked' : ''} oninput="configChangeHandler(this)">`;
        break;
      case 'dropdown':
        inputHtml = `<select id="${cfg.key}" oninput="configChangeHandler(this)">` +
          cfg.options.map(opt => `<option value="${opt}" ${cfg.value === opt ? 'selected' : ''}>${opt}</option>`).join('') +
          `</select>`;
        break;
      case 'number':
        inputHtml = `<input type="number" id="${cfg.key}" value="${cfg.value}" min="${cfg.min}" max="${cfg.max}" step="${cfg.step}" oninput="configChangeHandler(this)">`;
        break;
      case 'text':
        inputHtml = `<input type="text" id="${cfg.key}" value="${cfg.value}" oninput="configChangeHandler(this)">`;
        break;
      case 'hidden':
        continue;
    }

    content += `
      <div class="settings-row" ${cfg.showOnlyIf && !cfg.showOnlyIf() ? 'style="display:none"' : ''} 
      ${cfg.helptext ? `title="${cfg.helptext}"` : ''}>
        <label for="${cfg.key}">${cfg.display}</label>
        ${inputHtml}
      </div>`;
  }

  content += '</div>';
  return content;
}


let settingsWindow = null;
registerKeybinding('alt+k', () => {
  if (!settingsWindow) {
    settingsWindow = generateModalWindow({
      title: 'Settings',
      content: generateSettingsContent(),
      width: 400,
      height: 500,
      resizable: true,
    });
  } else {
    settingsWindow.remove();
    settingsWindow = null;
  }
});