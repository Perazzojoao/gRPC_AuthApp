import { IsNotEmpty, IsEmail, Matches } from 'class-validator';

export class CreateUserDto {
  @IsNotEmpty({ message: 'O campo email é obrigatório' })
  @IsEmail({}, { message: 'O campo email deve ser um email válido' })
  email: string;

  @IsNotEmpty({ message: 'O campo password é obrigatório' })
  @Matches(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*\W+).{6,30}$/, {
    message:
      'A senha deve conter pelo menos uma letra minúscula, uma letra maiúscula, um dígito, um caractere especial e ter entre 6 e 30 caracteres.',
  })
  password: string;
}
