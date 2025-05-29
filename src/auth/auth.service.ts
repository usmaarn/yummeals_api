import { BadRequestException, Injectable } from '@nestjs/common';
import { CreateUserDto } from './dtos/create-user.dto';
import { LoginDto } from './dtos/login.dto';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(private usersService: UsersService) {}
  register(dto: CreateUserDto) {}

  async login(dto: LoginDto) {
    const user = await this.usersService.findByEmail(dto.username);
    if (!user || !bcrypt.compareSync(dto.password, user.password)) {
      throw new BadRequestException({
        errors: { message: 'incorrect email or password' },
      });
    }
    return user;
  }
}
