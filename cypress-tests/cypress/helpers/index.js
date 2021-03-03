const email = () =>  Math.random().toString(36) + '@' + Math.random().toString(36)
const password = () => Math.random().toString(36)

export const gen = {
  email,
  password
}