export interface IAuthService {
  /**
   * Returns the backend URL to redirect the user to for GitHub OAuth.
   * Based on the Go backend router: `GET /auth/github`
   */
  getGithubAuthUrl(): string;

  // We could add other methods here later like logout, getSession, etc.
}
