<script setup lang="ts">
import { computed, ref } from 'vue'
import { useDraggable } from '../composables/useDraggable'

const { name = 'Window' } = defineProps<{ name?: string }>()

const width = ref(500)
const height = ref(400)
const { x, y, onMouseDown } = useDraggable(10, 10)

const style = computed(() => ({
  top: `${y.value}px`,
  left: `${x.value}px`,
  width: `${width.value}px`,
  height: `${height.value}px`,
}))
</script>

<template>
  <div class="bg-slate-100 absolute shadow-md overflow-y-hidden rounded-xl" :style="style">
    <header
      @mousedown="onMouseDown"
      class="p-2 bg-slate-50 cursor-grab active:cursor-grabbing select-none"
    >
      <p class="font-semibold">{{ name }}</p>
    </header>

    <div class="overflow-y-auto">
      <slot />
    </div>
  </div>
</template>
