import { ref } from 'vue'

export const useDraggable = (initialX = 0, initialY = 0) => {
  const x = ref(initialX)
  const y = ref(initialY)

  const isDragging = ref(false)
  const lastMousePosition = ref<{ x: number; y: number }>({ x: 0, y: 0 })

  const onMouseUp = () => {
    isDragging.value = false
    document.removeEventListener('mousemove', onMouseMove)
    document.removeEventListener('mouseup', onMouseUp)
  }

  const onMouseMove = (event: MouseEvent) => {
    if (!isDragging.value) return

    x.value += event.clientX - lastMousePosition.value.x
    y.value += event.clientY - lastMousePosition.value.y

    lastMousePosition.value = { x: event.clientX, y: event.clientY }
  }

  const onMouseDown = (event: MouseEvent) => {
    isDragging.value = true
    lastMousePosition.value = { x: event.clientX, y: event.clientY }
    document.addEventListener('mousemove', onMouseMove)
    document.addEventListener('mouseup', onMouseUp)
  }

  return { x, y, onMouseDown }
}
