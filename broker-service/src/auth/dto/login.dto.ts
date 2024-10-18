import { IsEmail, IsNotEmpty } from "class-validator";

export class LogInDto {

  @IsEmail({}, { message: 'O campo email deve ser um email válido'})
  @IsNotEmpty({ message: 'O campo email é obrigatório'})
  email: string;

  @IsNotEmpty({ message: 'O campo password é obrigatório'})
  password: string;
}