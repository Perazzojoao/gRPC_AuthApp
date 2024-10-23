import { Body, Controller, HttpCode, Inject, OnModuleInit, Post } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { LogInDto } from './dto/login.dto';
import { User, AuthService, AUTH_PACKAGE, USER_SERVICE_NAME, UserValidated, UserRequest, VerificationCodeRequest, UserResponse, ResendVerificationCodeRequest, ResendVerificationCodeResponse, SendResetPasswordRequest, SendResetPasswordResponse, ResetPasswordRequest, ResetPasswordResponse } from './interface/auth-service.interface';
import { Observable } from 'rxjs';
import { CreateUserDto } from './dto/create-user.dto';

@Controller('auth')
export class AuthController implements OnModuleInit {
  private authservice: AuthService;

  constructor(@Inject(AUTH_PACKAGE) private client: ClientGrpc) {}

  onModuleInit() {
    this.authservice = this.client.getService<AuthService>(USER_SERVICE_NAME);
  }

  @Post('register')
  @HttpCode(200)
  register(@Body() user: CreateUserDto): Observable<User> {
    return this.authservice.CreateUser(user);
  }

  @Post('login')
  @HttpCode(200)
  login(@Body() loginDto: LogInDto): Observable<UserValidated> {
    return this.authservice.ValidateUser(loginDto);
  }

  @Post('activate')
  @HttpCode(200)
  activateUser(@Body() user: VerificationCodeRequest): Observable<UserResponse> {
    return this.authservice.ActivateUser(user);
  }

  @Post('resend-code')
  @HttpCode(200)
  resendCode(@Body() user: ResendVerificationCodeRequest): Observable<ResendVerificationCodeResponse> {
    return this.authservice.ResendVerificationCode(user);
  }

  @Post('send-reset-password')
  @HttpCode(200)
  sendResetPassword(@Body() user: SendResetPasswordRequest): Observable<SendResetPasswordResponse> {
    return this.authservice.SendResetPassword(user);
  }

  @Post('reset-password')
  @HttpCode(200)
  resetPassword(@Body() user: ResetPasswordRequest): Observable<ResetPasswordResponse> {
    return this.authservice.ResetPassword(user);
  }
}
