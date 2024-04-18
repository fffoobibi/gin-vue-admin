/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  important: true,
  theme: {
    extend: {
      borderRadius: {
        "10": "10px"
      },
      backgroundColor: {
        "main": "#F5F5F5",
      }
    },
  },
  plugins: [],
  corePlugins: {
    preflight: false
  }
}

