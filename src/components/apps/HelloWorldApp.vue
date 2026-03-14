<script setup lang="ts">
import Window from '../Window.vue'
import type { Position, Size } from '../../types'

defineProps<{ id: string }>()

defineEmits<{ closed: [] }>()

const isMinimized = defineModel('isMinimized', { default: false })
const position = defineModel<Position>('position', { required: true })
const size = defineModel<Size>('size', { required: true })
</script>

<template>
  <Window
    name="Hello world"
    v-if="!isMinimized"
    v-model:position="position"
    v-model:size="size"
    @closed="$emit('closed')"
    @minimized="isMinimized = true"
  >
    <div class="flex w-full h-40 justify-center items-center flex-col">
      <p>Hello world!</p>
      <p>{{ id }}</p>
    </div>
  </Window>
</template>
