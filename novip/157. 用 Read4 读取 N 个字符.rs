// 157. 用 Read4 读取 N 个字符

/**
 * The read4 API is defined as.
 *     fn read4(&self,buf4: &mut [char]) -> i32;
 * You can call it using self.read4(buf4)
 */

impl Solution {
    pub fn read(&self, buf: &mut [char], n: i32) -> i32 {
        let mut buf4 = [0 as char; 4];
        let mut readn = 0;
        for i in 0..n {
            let modi = i%4;
            if modi == 0 {
                buf4.iter_mut().for_each(|x| *x=0 as char);
                readn = self.read4(&mut buf4);
            }
            if modi<readn {
                buf[i as usize] = buf4[modi as usize];
            } else {
                return i;
            }
        }
        n
    }
}

