<script setup lang="ts">
import Taskbar from './Taskbar.vue'
import Window from './Window.vue'
</script>

<template>
  <div
    class="absolute -z-10 inset-0 h-full w-full bg-[linear-gradient(to_right,#73737320_1px,transparent_1px),linear-gradient(to_bottom,#73737320_1px,transparent_1px)] bg-size-[20px_20px]"
  />

  <Window> hello world! </Window>
  <Taskbar />
</template>
