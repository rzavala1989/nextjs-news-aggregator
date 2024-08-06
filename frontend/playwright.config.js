// playwright.config.js
const { defineConfig } = require('@playwright/test');

module.exports = defineConfig({
  testDir: './tests',
  retries: 1,
  use: {
    baseURL: 'http://localhost:3000',
    browserName: 'chromium',
  },
});
