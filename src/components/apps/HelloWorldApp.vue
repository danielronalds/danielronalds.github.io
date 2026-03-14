<script setup lang="ts">
import Window from '../Window.vue'

defineProps<{ id: string }>()

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
    <div class="flex w-full h-40 justify-center items-center flex-col">
      <p>Hello world!</p>
      <p>{{ id }}</p>
    </div>
  </Window>
</template>
