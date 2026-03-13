import { ref } from 'vue'


export const useResizable = (
  initialWidth = 200,
  initialHeight = 200,
  minWidth: number = 100,
  minHeight: number = 100,
) => {
  const width = ref(initialWidth)
  const height = ref(initialHeight)

  const isResizing = ref(false)
  const lastMousePosition = ref({ x: 0, y: 0 })

  const onMouseUp = () => {
    isResizing.value = false
    document.removeEventListener('mousemove', onMouseMove)
    document.removeEventListener('mouseup', onMouseUp)
  }

  const onMouseMove = (event: MouseEvent) => {
    if (!isResizing.value) return

    width.value = Math.max(minWidth, width.value + event.clientX - lastMousePosition.value.x)
    height.value = Math.max(minHeight, height.value + event.clientY - lastMousePosition.value.y)

    lastMousePosition.value = { x: event.clientX, y: event.clientY }
  }

  const onMouseDown = (event: MouseEvent) => {
    isResizing.value = true
    lastMousePosition.value = { x: event.clientX, y: event.clientY }
    document.addEventListener('mousemove', onMouseMove)
    document.addEventListener('mouseup', onMouseUp)
  }

  return { width, height, onMouseDown }
}
