const email = () =>  Math.random().toString(36) + '@' + Math.random().toString(36)
const password = () => Math.random().toString(36)

export const MAIL_API = (Cypress.env('mail_url') || 'http://127.0.0.1:4437').replace(
  /\/$/,
  ''
)

export const gen = {
  email,
  password
}