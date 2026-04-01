import { container, IDIContainer } from "@/di/container";

/**
 * Registry Pattern.
 * Components/Use cases call getContainer() to retrieve the injected services.
 * This global singleton ensures we only instantiate our services once.
 */
export function getContainer(): IDIContainer {
  return container;
}
