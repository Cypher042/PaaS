import { GithubLoginUseCase } from "@/usecases/auth/GithubLoginUseCase";
import { getContainer } from "@/di/registry";
import { AuthPresenter } from "@/presenters/AuthPresenter";

/**
 * AuthController coordinates user interaction related to Authentication.
 * It is consumed by React Hooks.
 */
export class AuthController {
  private readonly githubLoginUC: GithubLoginUseCase;
  private readonly authPresenter: AuthPresenter;

  constructor() {
    const di = getContainer();

    // Instantiate use cases and pass the dependencies from the container
    this.githubLoginUC = new GithubLoginUseCase(di.authService);
    this.authPresenter = new AuthPresenter();
  }

  /**
   * Orchestrates the GitHub login flow
   * Returns the URL the client should navigate to
   */
  initiateGithubLogin(): string {
    // Could add zod validation here if taking inputs
    return this.githubLoginUC.execute();
  }
}
