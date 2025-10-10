<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import Header from './components/Header.vue'
import TopBar from './components/TopBar.vue';
</script>

<template>
  <Header />
  <HelloWorld />
  <TopBar />
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
