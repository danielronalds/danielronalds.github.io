<script setup lang="ts">
import { computed } from 'vue'
import { useDraggable } from '../composables/useDraggable'
import { useResizable } from '../composables/useResizable'

import { Minus, X } from 'lucide-vue-next'

const { name = 'Window' } = defineProps<{ name?: string }>()

defineEmits<{ closed: []; minimized: [] }>()

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
      class="p-2 bg-base-100 cursor-grab active:cursor-grabbing select-none flex flex-row justify-between"
    >
      <p class="font-semibold">{{ name }}</p>

      <div class="flex flex-row gap-2">
        <button @click="$emit('minimized')" class="cursor-pointer">
          <minus :size="18" />
        </button>
        <button @click="$emit('closed')" class="cursor-pointer">
          <X :size="18" />
        </button>
      </div>
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
