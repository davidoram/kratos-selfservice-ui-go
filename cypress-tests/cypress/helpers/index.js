const email = () =>  Math.random().toString(36) + '@' + Math.random().toString(36)
const password = () => Math.random().toString(36)

export const APP_URL = (Cypress.env('app_url') || 'http://127.0.0.1:4455').replace(
  /\/$/,
  ''
)
export const MOBILE_URL = (Cypress.env('mobile_url') || 'http://127.0.0.1:4457').replace(
  /\/$/,
  ''
)
export const KRATOS_ADMIN = (Cypress.env('kratos_admin') || 'http://127.0.0.1:4434')
  .replace()
  .replace(/\/$/, '')
export const KRATOS_PUBLIC = (Cypress.env('kratos_public') || 'http://127.0.0.1:4433')
  .replace()
  .replace(/\/$/, '')
export const MAIL_API = (Cypress.env('mail_url') || 'http://127.0.0.1:4437').replace(
  /\/$/,
  ''
)

export const gen = {
  email,
  password,
  identity: () => ({ email: email(), password: password() })
}