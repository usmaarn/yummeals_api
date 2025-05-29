import { Body, Controller, Get, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { LoginDto } from './dtos/login.dto';

@Controller('v1/auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Get('user')
  getAuthenticatedUser(): any {
    return null;
  }

  @Post('register')
  register() {}

  @Post('login')
  login(@Body() dto: LoginDto) {
    return this.authService.login(dto);
  }

  @Post('logout')
  logout() {}

  @Post('password/sent-otp')
  sendPasswordResetOtp() {}

  @Post('password/reset-password')
  resetPassword() {}
}
