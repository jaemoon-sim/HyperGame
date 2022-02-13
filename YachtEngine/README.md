## build
```
go build --o yachtengine main.go
```

## run
```
./yachtengine --output log.json asdf,9910 zxcv,9990
```

### params
|key|required|설명|예시|
|---|---|---|---|
|output|N|yachtengine 실행 후 전체 로그 출력 path. 비어있을 경우 stdout|`--output ./log.json`|
|(없음)|Y|`이름,port번호` 형식. 여러 명 참여 가능. 이름 중복 가능|`jb,1234 jb2,1234 jb3,1234`|

## output

|key|설명|
|---|---
|log|player의 decision이 발생하기 직전의 state와 player의 decision을 묶어서 json array로 가지고 있음
|`log > state`|player의 decision이 발생하기 직전의 state
|`log > decision`| `log > state`에 대한 player의 decision
|`final`| 다 끝난 뒤 게임 state|

```
{
 "log": [
  {
   "state": {
    "turn": 1,
    "player": "jb",
    "trial": 1,
    "dices": [
     5,
     3,
     1,
     3,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": null,
      "subtotal": 0,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 0
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "threes"
   }
  },
  {
   "state": {
    "turn": 2,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     1,
     5,
     1,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": null,
      "fives": null,
      "sixes": null,
      "subtotal": 6,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 6
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "sixes"
   }
  },
  {
   "state": {
    "turn": 3,
    "player": "jb",
    "trial": 1,
    "dices": [
     4,
     1,
     2,
     2,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": null,
      "fives": null,
      "sixes": 6,
      "subtotal": 12,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 12
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fives"
   }
  },
  {
   "state": {
    "turn": 4,
    "player": "jb",
    "trial": 1,
    "dices": [
     3,
     1,
     4,
     5,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": null,
      "fives": 5,
      "sixes": 6,
      "subtotal": 17,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 17
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "smallStraight"
   }
  },
  {
   "state": {
    "turn": 5,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     1,
     1,
     2,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": null,
      "fives": 5,
      "sixes": 6,
      "subtotal": 17,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 32
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fullHouse"
   }
  },
  {
   "state": {
    "turn": 6,
    "player": "jb",
    "trial": 1,
    "dices": [
     5,
     5,
     3,
     4,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": null,
      "fives": 5,
      "sixes": 6,
      "subtotal": 17,
      "choice": null,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 39
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "choice"
   }
  },
  {
   "state": {
    "turn": 7,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     4,
     6,
     1,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": null,
      "fives": 5,
      "sixes": 6,
      "subtotal": 17,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 58
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fours"
   }
  },
  {
   "state": {
    "turn": 8,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     2,
     2,
     2,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": 4,
      "fives": 5,
      "sixes": 6,
      "subtotal": 21,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 62
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "deuces"
   }
  },
  {
   "state": {
    "turn": 9,
    "player": "jb",
    "trial": 1,
    "dices": [
     3,
     3,
     5,
     6,
     1
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 6,
      "threes": 6,
      "fours": 4,
      "fives": 5,
      "sixes": 6,
      "subtotal": 27,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 68
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "aces"
   }
  },
  {
   "state": {
    "turn": 10,
    "player": "jb",
    "trial": 1,
    "dices": [
     2,
     4,
     5,
     2,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 6,
      "threes": 6,
      "fours": 4,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 69
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "largeStraight"
   }
  },
  {
   "state": {
    "turn": 11,
    "player": "jb",
    "trial": 1,
    "dices": [
     3,
     5,
     4,
     6,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 6,
      "threes": 6,
      "fours": 4,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": 0,
      "yacht": null,
      "total": 69
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "yacht"
   }
  },
  {
   "state": {
    "turn": 12,
    "player": "jb",
    "trial": 1,
    "dices": [
     4,
     6,
     3,
     4,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 6,
      "threes": 6,
      "fours": 4,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 7,
      "smallStraight": 15,
      "largeStraight": 0,
      "yacht": 0,
      "total": 69
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fourKind"
   }
  }
 ],
 "final": {
  "dices": [
   0,
   0,
   0,
   0,
   0
  ],
  "scoreBoard": {
   "jb": {
    "aces": 1,
    "deuces": 6,
    "threes": 6,
    "fours": 4,
    "fives": 5,
    "sixes": 6,
    "subtotal": 28,
    "choice": 19,
    "fourKind": 0,
    "fullHouse": 7,
    "smallStraight": 15,
    "largeStraight": 0,
    "yacht": 0,
    "total": 69
   }
  }
 }
}
```
