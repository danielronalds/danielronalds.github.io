<script setup lang="ts">
import Window from '../Window.vue'

defineEmits<{ closed: [] }>()

const isMinimized = defineModel('isMinimized', { default: false })
</script>

<template>
  <Window
    name="Hello world"
    v-if="!isMinimized"
    @closed="$emit('closed')"
    @minimized="isMinimized = true"
  >
    <div class="flex w-full h-40 justify-center items-center">
      <p>Hello world!</p>
    </div>
  </Window>
</template>
