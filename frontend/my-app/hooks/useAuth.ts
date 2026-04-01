import { useMemo, useState } from "react";
import { AuthController } from "@/controllers/AuthController";

export function useAuth() {
  // Normally we would use React Query here for async state, but doing simple state for now
  const [isLoading, setIsLoading] = useState(false);
  const authController = useMemo(() => new AuthController(), []);

  const loginWithGithub = () => {
    setIsLoading(true);
    // Get the redirect URL from the controller
    const url = authController.initiateGithubLogin();
    // Redirect the browser
    window.location.href = url;
  };

  return {
    loginWithGithub,
    isLoading,
  };
}
