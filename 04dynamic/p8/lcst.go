package p8

func Lcst2(s1, s2 string) string {
	row, col := 0, len(s2)
	maxLen, end := 0, 0
	for row < len(s1) {
		i, j := row, col
		l := 0
		for i < len(s1) && j < len(s2) {
			if s1[i] != s2[j] {
				l = 0
			} else {
				l++
			}
			if l > maxLen {
				maxLen = l
				end = i
			}
			i++
			j++
		}
		if col > 0 {
			col--
		} else {
			row++
		}
	}
	return s1[end-maxLen+1 : end+1]
}

/**

public String lcst2(String str1, String str2) {
              if (str1 == null || str2 == null || str1.equals("") || str2.equals("")) {
                  return "";
              }
              char[] chs1 = str1.toCharArray();
              char[] chs2 = str2.toCharArray();
              int row = 0; // 斜线开始位置的行
              int col = chs2.length - 1; // 斜线开始位置的列
              int max = 0; // 记录最大长度
              int end = 0; // 最大长度更新时，记录子串的结尾位置
              while (row < chs1.length) {
                  int i = row;
                  int j = col;
                  int len = 0;
                  // 从(i, j)开始向右下方遍历
                  while (i < chs1.length && j < chs2.length) {
                          if (chs1[i] ! = chs2[j]) {
                                  len = 0;
                          } else {
                                  len++;
                          }
                          // 记录最大值，以及结束字符的位置
                          if (len > max) {
                                  end = i;
                                  max = len;
                          }
                          i++;
                          j++;
                  }
                  if (col > 0) { // 斜线开始位置的列先向左移动
                          col--;
                  } else { // 列移动到最左之后，行向下移动
                          row++;
                  }
              }
              return str1.substring(end - max + 1, end + 1);
          }
*/
