/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'Anta': ['Anta', 'sans-serif'],
        'Roboto': ['Roboto', 'sans-serif'],
      },
      boxShadow: {
        'equal-md': '0px 0px 8px 2px rgba(0, 0, 0, 0.1)',
      }
    },
  },
  plugins: [],
}

