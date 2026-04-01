import { User } from "@/types/auth.types";

export interface AuthUIModel {
  displayName: string;
  avatarUrl?: string; // We can extract this or pass it later
  isAuthenticated: boolean;
}

export class AuthPresenter {
  /**
   * Transforms raw backend user model into a safe UI format
   */
  toUIModel(user: User | null): AuthUIModel {
    if (!user) {
      return {
        displayName: "Guest",
        isAuthenticated: false,
      };
    }

    return {
      displayName: user.username || user.githubUsername || "User",
      isAuthenticated: true,
      // If we grab avatars, we'd add it here
    };
  }
}
