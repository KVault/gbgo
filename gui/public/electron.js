const Menu  = require("electron-create-menu")
const { app, BrowserWindow } = require('electron');
const path = require("path");
const isDev = require("electron-is-dev");

/**
 * Holds a reference to all the windows that the app can have 
 */
let windows = {
    mainWindow: null,
    memoryViewerWindow: null,
}

/**
 * Creates the main window, this function will save the window reference in the "windows" variable
 */
function createMainWindow() {
    windows.mainWindow = new BrowserWindow({ width: 160, height: 144 });
    windows.mainWindow.loadURL(
        isDev
            ? "http://localhost:3000/game"
            : `file://${path.join(__dirname, "../build/index.html/game")}`
    );
    windows.mainWindow.on("closed", () => (windows.mainWindow = null));
    setMainMenu();
}

/**
 * Creates the memory viewer window
 */
function createMemoryViewerWindow(){
    windows.memoryViewerWindow = new BrowserWindow({ width: 160, height: 300 });
    windows.memoryViewerWindow.loadURL(
        isDev
            ? "http://localhost:3000/memoryViewer"
            : `file://${path.join(__dirname, "../build/index.html/memoryViewer")}`
    );
    windows.memoryViewerWindow.on("closed", () => (windows.memoryViewerWindow = null));
}


app.on("ready", createMainWindow);
app.on("window-all-closed", () => {
    if (process.platform !== "darwin") {
        app.quit();
    }
});
app.on("activate", () => {
    if (windows.mainWindow === null) {
        createMainWindow();
    }
});

function setMainMenu() {
    Menu();

    Menu((defaultMenu, separator) => {
 
        defaultMenu.push({
          label: 'Debugger',
          submenu: [
            {label: 'Memory Viewer', click:  () =>createMemoryViewerWindow()},
            separator(),
            {label: 'my second item'},
          ],
        });
       
        return defaultMenu;
      });
}