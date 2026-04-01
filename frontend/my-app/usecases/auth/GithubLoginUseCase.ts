import { IAuthService } from "@/interfaces/IAuthService";

/**
 * Single job: Get the URL needed to start the Github OAuth flow
 */
export class GithubLoginUseCase {
  constructor(private readonly authService: IAuthService) {}

  execute(): string {
    return this.authService.getGithubAuthUrl();
  }
}
