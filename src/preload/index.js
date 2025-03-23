import { contextBridge, ipcRenderer } from 'electron'
import { electronAPI } from '@electron-toolkit/preload'

const versions = {
  vue: '3.5.13',
  golang: '1.24.1',
  bluebell: '1.0.0'
}

// Custom APIs for renderer
const api = {
  versions,
  close: () => ipcRenderer.send('window-close'),
  minimize: () => ipcRenderer.send('window-minimize'),
  maximize: () => ipcRenderer.send('window-maximize'),
  onWindowMaximizeUnmaximize: (callback) => ipcRenderer.on('window-maximize-unmaximize', (_, isMaximized) => callback(isMaximized)),
  onWindowBlur: (callback) => ipcRenderer.on('window-blur', callback),
  openNewWindow: (routeName, param) => ipcRenderer.send('open-new-window', routeName, param),
  selectFolder: () => ipcRenderer.invoke('select-folder'),
  openFile: (path) => ipcRenderer.invoke('open-file', path),
  openInExplorer: (path) => ipcRenderer.invoke('open-in-explorer', path),
  getConfig: (key) => ipcRenderer.invoke('get-config', key),
  setConfig: (key, val) => ipcRenderer.invoke('set-config', key, val)
}

// Use `contextBridge` APIs to expose Electron APIs to
// renderer only if context isolation is enabled, otherwise
// just add to the DOM global.
if (process.contextIsolated) {
  try {
    contextBridge.exposeInMainWorld('electron', electronAPI)
    contextBridge.exposeInMainWorld('api', api)
  } catch (error) {
    console.error(error)
  }
} else {
  window.electron = electronAPI
  window.api = api
}
