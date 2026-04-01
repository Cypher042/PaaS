// Types corresponding to the Go backend User model

export interface User {
  id: string;
  username: string;
  githubUsername?: string;
  email?: string;
  githubToken?: string;
}

export interface AuthResponse {
  message: string;
  user: User;
}
