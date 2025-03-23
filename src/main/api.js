import { BrowserWindow, ipcMain, dialog, shell } from 'electron'
import { getConfig, setConfig } from './store'
import { exec } from 'child_process'

ipcMain.handle('select-folder', async () => {
  const window = BrowserWindow.getFocusedWindow()
  const result = await dialog.showOpenDialog(window, {
    properties: ['openDirectory']
  })

  if (!result.canceled && result.filePaths.length > 0) {
    return result.filePaths[0]
  }
  return null
})

ipcMain.handle('open-file', async (event, path) => {
  return await shell.openPath(path)
})

ipcMain.handle('open-in-explorer', (event, path) => {
  let command = `explorer.exe /select,${path}`
  if (process.platform === "linux") {
    command = `xdg-open "${path}"`
  } else if (process.platform === "darwin") {
    command = `open "${path}"`
  }
  exec(command,(error, stdout, stderr) =>{
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
})

ipcMain.handle('get-config', (event, key) => {
  return getConfig(key)
})

ipcMain.handle('set-config', (event, key, value) => {
  setConfig(key, value)
  return true
})