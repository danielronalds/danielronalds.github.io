import type { Position, Size } from './types'

export const getDefaultWindowPosition = (size: Size): Position => ({
  x: window.innerWidth / 2 - size.width / 2,
  y: window.innerHeight / 2 - size.height / 2,
})

export const generateId = () => {
  const alphabet = 'abcdefghijklmnopqrstuvwxyz'
  let randomString = ''

  for (let i = 0; i < 26; i++) {
    randomString += alphabet[Math.floor(Math.random() * alphabet.length)]
  }

  return randomString
}

