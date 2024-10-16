import { config } from 'dotenv';
import { ReflectionService } from '@grpc/reflection';
import { ClientOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

config();

export const grpcClientOptions: ClientOptions = {
  transport: Transport.GRPC,
  options: {
    package: 'proto',
    url: process.env.AUTH_SERVICE_URL || 'localhost:8000',
    protoPath: join(__dirname, './auth/auth.proto'),
    onLoadPackageDefinition: (pkg, server) => {
      new ReflectionService(pkg).addToServer(server);
    },
  },
};