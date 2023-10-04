/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin')
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    container: {
      center: true,
      padding: '1rem'
    },
    boxShadow: {
      'normal': '0px 0px 4px 0px rgba(0, 0, 0, 0.14)',
    },
    extend: {
      colors: {
        'main-bg-web': '#eee',
        'main-blue-web': '#1875FF',
        'main-dark-text-web': '#394149',
        'main-gray-text-web': '#726E6E',
        'second-gray-text-web': '#595757',
        'third-gray-text-web': '#D9D9D9',
        'main-green-web': '#00886B',
        'main-red-web': '#E7373C',
        'main-yellow-web': '#E9A600',
        'second-yellow-web': '#FFEBD3',
        'main-violet-web': '#6328DF',
        'main-light-web': '#fff',

      },
      fontFamily: {
        'Ray-Medium': 'Ray-Medium',
        'Ray-Light': 'Ray-Light',
        'Ray-ExtraBold': 'Ray-ExtraBold',
        'Ray-ExtraBlack': 'Ray-ExtraBlack',
        'Ray-Bold': 'Ray-Bold',
        'Ray-Black': 'Ray-Black',

        'ANJOMANFANUM-HEAVY': 'ANJOMANFANUM-HEAVY',
        'ANJOMANFANUM-MEDIUM': 'ANJOMANFANUM-MEDIUM',
        'ANJOMANFANUM-SEMIBOLD': 'ANJOMANFANUM-SEMIBOLD',
        'ANJOMANFANUM-THIN': 'ANJOMANFANUM-THIN',
        'ANJOMANFANUM-ULTRABOLD': 'ANJOMANFANUM-ULTRABOLD'
      }
    },
  },
  plugins: [
    plugin(function ({ addVariant }) {
      addVariant("child", "& > *")
      addVariant("child-hover", "& > *:hover")
    })
  ],
}

