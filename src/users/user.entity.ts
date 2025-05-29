import { UserType } from 'src/lib/enums';
import { Token } from 'src/tokens/token.entity';
import { Column, Entity, OneToMany, PrimaryGeneratedColumn } from 'typeorm';

@Entity({ name: 'users' })
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ nullable: false })
  name: string;

  @Column({ unique: true })
  email: string;

  @Column()
  phoneNumber: string;

  @Column()
  password: string;

  @Column({ default: UserType.USER })
  type: UserType;

  @Column({ default: true })
  isActive: boolean;

  @OneToMany(() => Token, (token) => token.user)
  tokens: Token[];
}
