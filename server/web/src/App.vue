<script setup>
  import NavBar from './components/NavBar.vue'
  import Footer from './components/Footer.vue'
  import { RouterView } from 'vue-router'
</script>

<template>
  <NavBar />

  <main>
    <RouterView />
  </main>
  
  <Footer />
</template>

<style scoped>

</style>
