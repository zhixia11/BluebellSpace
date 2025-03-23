import { app, shell, BrowserWindow, ipcMain, protocol } from 'electron'
import { join } from 'path'
import { electronApp, optimizer, is } from '@electron-toolkit/utils'
import './api'
import icon from '../../resources/icon.png?asset'
import fs from 'fs'
import { exec } from 'child_process'

const enginePath = join(__dirname, is.dev ? '../../resources/engine/engine.exe' : 'resources/engine/engine.exe');

protocol.registerSchemesAsPrivileged([
  {
    scheme: 'local', // 协议名称
    privileges: {
      standard: true,
      secure: true,
      supportFetchAPI: true,
      allowServiceWorkers: true,
      corsEnabled: true,
      stream: true,
    },
  },
])

function createWindow(routeName, width, height) {
  // Create the browser window.
  const window = new BrowserWindow({
    width: width || 820,
    height: height || 660,
    show: false,
    frame: false,
    resizable: false,
    // transparent: true,
    autoHideMenuBar: true,
    ...(process.platform === 'linux' ? { icon } : {}),
    webPreferences: {
      preload: join(__dirname, '../preload/index.mjs'),
      sandbox: false
    }
  })
  window.removeMenu()

  window.on('ready-to-show', () => {
    window.show()
  })

  window.on('blur', () => {
    window.webContents.send('window-blur');
  })

  window.on('maximize', () => {
    window.webContents.send('window-maximize-unmaximize', true)
  })
  window.on('unmaximize', () => {
    window.webContents.send('window-maximize-unmaximize', false)
  })

  window.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url)
    return { action: 'deny' }
  })

  // HMR for renderer base on electron-vite cli.
  // Load the remote URL for development or the local html file for production.
  if (!!routeName) {
    routeName = '#' + routeName
  } else {
    routeName = ''
  }
  if (is.dev && process.env['ELECTRON_RENDERER_URL']) {
    window.loadURL(process.env['ELECTRON_RENDERER_URL'] + routeName)
  } else {
    window.loadFile(join(__dirname, '../renderer/index.html' + routeName))
  }
  return window
}

// 主窗口
let mainWindow = null

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {
  // Set app user model id for windows
  electronApp.setAppUserModelId('com.electron')

  // Default open or close DevTools by F12 in development
  // and ignore CommandOrControl + R in production.
  // see https://github.com/alex8088/electron-toolkit/tree/master/packages/utils
  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window)
  })

  // IPC test
  ipcMain.on('ping', () => console.log('pong'))

  mainWindow = createWindow()

  console.log('exec -> ', enginePath)

  // 启动engine
  if (!is.dev) {
    exec(enginePath, { encoding: 'utf8' }, (error, stdout, stderr) => {
      if (error) {
        console.error(`Error: ${error.message}`);
        return;
      }
      if (stderr) {
        console.error(`Stderr: ${stderr}`);
        return;
      }
      console.log(`Stdout: ${stdout}`);
    })
  }

  app.on('activate', function () {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })


  protocol.handle("local", async (request) => {
    let url = decodeURIComponent(request.url)
    url = url.slice(8)
    const path = convertPath(url)
    const data = fs.readFileSync(path)
    return new Response(data)
  })
})

function convertPath(src) {
  let path = src
  if (process.platform === "win32") {
    const block = src[0]
    path = block + ":" + src.slice(1)
  }
  return path
}

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.
// 窗口关闭
ipcMain.on('window-close', () => {
  const window = BrowserWindow.getFocusedWindow()
  //如果是主窗口就退出应用
  if (window === mainWindow) {
    app.quit()
  } else {
    window?.close()
  }
})
// 窗口最小化
ipcMain.on('window-minimize', () => {
  const window = BrowserWindow.getFocusedWindow()
  window?.minimize()
})
// 窗口最大化和还原
ipcMain.on('window-maximize', () => {
  const window = BrowserWindow.getFocusedWindow()
  if (window?.isMaximized()) {
    window?.unmaximize()
  } else {
    window?.maximize()
  }
})
// 窗口管理
const windowMap = {}
// 打开新窗口
ipcMain.on('open-new-window', (event, routeName) => {
  if (windowMap[routeName]) {
    windowMap[routeName].show()
    return
  }
  const window = createWindow(routeName)
  windowMap[routeName] = window
})
