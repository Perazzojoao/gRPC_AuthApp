import { Module } from '@nestjs/common';
import { ClientsModule } from '@nestjs/microservices';
import { grpcClientOptions } from 'src/grpc-client.options';
import { AuthController } from './auth.controller';
import { AUTH_PACKAGE } from './interface/auth-service.interface';

@Module({
  imports: [
    ClientsModule.register([
      {
        name:AUTH_PACKAGE,
        ...grpcClientOptions,
      },
    ]),
  ],
  controllers: [AuthController],
  providers: [],
})
export class AuthModule {}
