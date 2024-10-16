import { Body, Controller, Inject, OnModuleInit, Post } from '@nestjs/common';
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
  register(@Body() user: CreateUserDto): Observable<User> {
    return this.authservice.CreateUser(user);
  }

  @Post('login')
  login(@Body() loginDto: LogInDto): Observable<UserValidated> {
    return this.authservice.ValidateUser(loginDto);
  }

  @Post('activate')
  activateUser(@Body() user: VerificationCodeRequest): Observable<UserResponse> {
    return this.authservice.ActivateUser(user);
  }

  @Post('resend-code')
  resendCode(@Body() user: ResendVerificationCodeRequest): Observable<ResendVerificationCodeResponse> {
    return this.authservice.ResendVerificationCode(user);
  }

  @Post('send-reset-password')
  sendResetPassword(@Body() user: SendResetPasswordRequest): Observable<SendResetPasswordResponse> {
    return this.authservice.SendResetPassword(user);
  }

  @Post('reset-password')
  resetPassword(@Body() user: ResetPasswordRequest): Observable<ResetPasswordResponse> {
    return this.authservice.ResetPassword(user);
  }
}
