## build
```
go run main.go --o yachtengine
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
     1,
     2,
     3,
     1,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": null,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null
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
    "turn": 2,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     5,
     3,
     2,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": 5,
      "sixes": null,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null
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
    "turn": 3,
    "player": "jb",
    "trial": 1,
    "dices": [
     4,
     6,
     3,
     4,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": 5,
      "sixes": null,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
    "turn": 4,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     1,
     5,
     5,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": null,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
    "turn": 5,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     4,
     6,
     4,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
    "turn": 6,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     3,
     1,
     5,
     6
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
    "turn": 7,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     2,
     6,
     5,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
    "turn": 8,
    "player": "jb",
    "trial": 1,
    "dices": [
     4,
     5,
     5,
     4,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 2,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
    "turn": 9,
    "player": "jb",
    "trial": 1,
    "dices": [
     3,
     5,
     3,
     1,
     1
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 2,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
     4,
     2,
     3,
     4,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": 2,
      "deuces": 2,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null
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
     2,
     2,
     5,
     5,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": 2,
      "deuces": 2,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": 0,
      "yacht": null
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
     2,
     6,
     1,
     3,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": 2,
      "deuces": 2,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "choice": 23,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": 0,
      "yacht": 0
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
  "scoreBoard": {
   "jb": {
    "aces": 2,
    "deuces": 2,
    "threes": 3,
    "fours": 8,
    "fives": 5,
    "sixes": 6,
    "choice": 23,
    "fourKind": 0,
    "fullHouse": 0,
    "smallStraight": 15,
    "largeStraight": 0,
    "yacht": 0
   }
  }
 }
}
```
