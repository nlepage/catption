importScripts('https://cdn.jsdelivr.net/gh/nlepage/go-wasm-http-server@574bb8928c6a7a5c4feeaf7cee9530ea5e3c04b1/sw.js')

addEventListener('install', (event) => {
  event.waitUntil(skipWaiting())
})
  
addEventListener('activate', event => {
  event.waitUntil(clients.claim())
})

registerWasmHTTPListener('catption.wasm', undefined, ['http'])
