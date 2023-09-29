/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors:{
        'main-blue-web' : '#1875FF',
        'main-dark-text-web' : '#394149'
        
      }
    },
  },
  plugins: [],
}

