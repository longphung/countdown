export const createRipple = <T extends HTMLElement> (
  event: React.MouseEvent<T>
): void => {
  const button: HTMLElement = event.currentTarget;
  // Calculate circle parameters
  const buttonCoordinates = button.getBoundingClientRect();
  const circle = document.createElement("span");
  const diameter = Math.max(button.clientWidth, button.clientHeight);
  const radius = diameter / 2;
  // Get circle style
  circle.style.width = `${diameter}px`;
  circle.style.height = `${diameter}px`;
  circle.style.left = `${event.clientX - buttonCoordinates.x - radius}px`;
  circle.style.top = `${event.clientY - buttonCoordinates.y - radius}px`;
  circle.classList.add("ripple");
  // Handle DOM
  const ripple: Element = button.getElementsByClassName("ripple")[0];
  // Remove the previous ripple
  if (ripple && ripple.parentElement) {
    ripple.parentElement.removeChild(ripple);
  }
  button.appendChild(circle);
};
