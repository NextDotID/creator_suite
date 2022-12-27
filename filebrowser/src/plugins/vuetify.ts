// import { getCurrentInstance } from 'vue'

// Vuetify
import * as directives from 'vuetify/directives'
import * as components from 'vuetify/components'
import { createVuetify, type ThemeDefinition } from 'vuetify'

// Styles
import 'vuetify/styles'
// import { aliases, mdi } from 'vuetify/iconsets/mdi-svg'
// import { fa } from 'vuetify/iconsets/fa'
import { aliases, fa } from 'vuetify/iconsets/fa'
import { mdi } from 'vuetify/lib/iconsets/mdi'

// make sure to also import the coresponding css
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import '@fortawesome/fontawesome-free/css/all.css' // Ensure your project is capable of handling css files

// Theme
const light: ThemeDefinition = {
  dark: false,
  colors: {
    background: '#efefef',
    surface: '#fff',
    primary: '#7743DB',
    'primary-darken-1': '#3B3486',
    secondary: '#ef476f',
    'secondary-darken-1': '#ed315d',
    error: '#B00020',
    info: '#4496F3',
    success: '#16DB93',
    warning: '#CB8C00',
  },
}

const dark: ThemeDefinition = {
  dark: true,
  colors: {
    background: '#181833',
    surface: '#161644',
    primary: '#7743DB',
    'primary-darken-1': '#3B3486',
    secondary: '#ef476f',
    'secondary-darken-1': '#ed315d',
    error: '#B00020',
    info: '#4496F3',
    success: '#16DB93',
    warning: '#CB8C00',
  },
}

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'dark',
    themes: {
      light,
      dark,
    },
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
      fa,
    },
  },
})

export default vuetify
