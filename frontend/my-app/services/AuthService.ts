import { api } from "@/lib/api";
import { IAuthService } from "@/interfaces/IAuthService";

export class AuthService implements IAuthService {
  getGithubAuthUrl(): string {
    // The backend handles the complete OAuth flow and redirecting back to Github.
    // So the frontend just needs to redirect the user to this backend URL.
    // The base URL is handled by api.defaults.baseURL but since this is a redirect
    // we need the full absolute URL string.
    
    const baseUrl = api.defaults.baseURL || "http://localhost:8080";
    return `${baseUrl}/auth/github`;
  }
}
