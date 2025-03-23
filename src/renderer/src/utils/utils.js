
export const getFileName = (path) => {
  // 根据是否是Windows判断
  let char = '/'
  if (isWindows()) {
    char = '\\'
  }
  const arr = path.split(char)
  return arr[arr.length - 1]
}

export const formatSize = (size) => {
  if (size < 1024) {
    return size + 'B'
  } else if (size < 1024 * 1024) {
    return (size / 1024).toFixed(2) + 'KB'
  } else if (size < 1024 * 1024 * 1024) {
    return (size / (1024 * 1024)).toFixed(2) + 'MB'
  }
}

const isWindows = () => {
  return window.navigator.userAgent.indexOf('Win') !== -1
}

export const getExifDisplayName = (key) => {
  switch (key) {
    case 'Size':
      return '大小'
    case 'ModTime':
      return '修改时间'
  }
  return ''
}