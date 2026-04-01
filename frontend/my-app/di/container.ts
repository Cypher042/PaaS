import { IAuthService } from "@/interfaces/IAuthService";
import { AuthService } from "@/services/AuthService";

/**
 * The DI container holds instances of all services.
 * We instantiate them here.
 */
export interface IDIContainer {
  authService: IAuthService;
  // added more services here later e.g., projectService: IProjectService
}

// Instantiate the concrete services
const authService = new AuthService();

// Register them in the container
export const container: IDIContainer = {
  authService,
};
