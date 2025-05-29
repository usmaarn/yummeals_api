import { User } from 'src/users/user.entity';
import { Column, Entity, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';

@Entity({ name: 'tokens' })
export class Token {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne(() => User, (user) => user.tokens)
  user: User;

  @Column({ unique: true })
  accessToken: string;

  @Column({ unique: true })
  refreshToken: string;

  @Column({ type: 'timestamptz' })
  accessTokenExpiresAt: Date;

  @Column({ type: 'timestamptz' })
  refreshTokenExpiresAt: Date;
}
