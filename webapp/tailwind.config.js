module.exports = {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    fontFamily: {
      sans: ["Montserrat", "sans-serif"],
      serif: ["Georgia", "serif"],
      mono: ["Menlo", "monospace"],
      cursive: ["Oleo Script Swash Caps", "cursive"],
    },
    extend: {
      colors: {
        rippleWhite: "rgba(255, 255, 255, 0.7)",
      },
      keyframes: {
        ripple: {
          to: {
            transform: "scale(4)",
            opacity: "0",
          },
        },
      },
      animation: {
        ripple: "ripple 600ms linear",
      },
    },
  },
  plugins: [require("daisyui")],
};
