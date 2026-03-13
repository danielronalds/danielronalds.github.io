<script setup lang="ts">
import { computed } from 'vue'
import { useDraggable } from '../composables/useDraggable'
import { useResizable } from '../composables/useResizable'

const { name = 'Window' } = defineProps<{ name?: string }>()

const { x, y, onMouseDown: onDragMouseDown } = useDraggable(10, 10)
const { width, height, onMouseDown: onResizeMouseDown, isResizing } = useResizable(500, 400)

const style = computed(() => ({
  top: `${y.value}px`,
  left: `${x.value}px`,
  width: `${width.value}px`,
  height: `${height.value}px`,
}))
</script>

<template>
  <div class="bg-base-200 absolute shadow-lg overflow-y-hidden rounded-xl" :style="style">
    <header
      @mousedown="onDragMouseDown"
      class="p-2 bg-base-100 cursor-grab active:cursor-grabbing select-none"
    >
      <p class="font-semibold">{{ name }}</p>
    </header>

    <div class="overflow-y-auto" :class="{ 'select-none': isResizing }">
      <slot />
    </div>

    <div
      @mousedown="onResizeMouseDown"
      class="absolute bottom-0 right-0 w-3 h-3 cursor-nwse-resize"
    />
  </div>
</template>
