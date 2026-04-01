import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

export function middleware(request: NextRequest) {
  // If the user is trying to access the dashboard
  if (request.nextUrl.pathname.startsWith('/dashboard')) {
    const jwt = request.cookies.get('jwt')
    
    // If no JWT found, redirect to login
    if (!jwt) {
      return NextResponse.redirect(new URL('/login', request.url))
    }
  }
  
  // Protect other routes here like /projects or /settings
  
  return NextResponse.next()
}

// See "Matching Paths" below to learn more
export const config = {
  matcher: [
    '/dashboard/:path*',
    '/projects/:path*',
  ],
}
