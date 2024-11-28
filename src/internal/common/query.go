package common

type QueryPara struct {
	IdIn        []int                         `json:"idIn,omitempty"`        // 实体ID列表
	Name        string                        `json:"name,omitempty"`        // 实体名称
	NameLike    string                        `json:"nameLike,omitempty"`    // 实体名称，前精确，后模糊
	Code        string                        `json:"code,omitempty"`        // 编码
	CodeLike    string                        `json:"codeLike,omitempty"`    // 编码，前精确，后模糊
	ParentId    int                           `json:"parentId,omitempty"`    // 查直接子实例
	AllSubFrom  string                        `json:"allSubFrom,omitempty"`  // 查所有直接和间接的子实例（不包含自己）
	Label       map[string]string             `json:"label,omitempty"`       // 键和值都需要匹配
	NLabel      map[string]int                `json:"nLabel,omitempty"`      // 键和值都需要匹配,数值型标签
	LabelLike   map[string]string             `json:"labelLike,omitempty"`   // label的值可以进行前精确后模糊查询
	LabelKey    []string                      `json:"labelKey,omitempty"`    // 满足指定的 key 即可
	NLabelKey   []string                      `json:"nLabelKey,omitempty"`   // 满足指定的 key 即可,数值型标签
	LabelIn     map[string][]string           `json:"labelIn,omitempty"`     // 一个 key 中的 value 值只要有一个满足就可以。
	NLabelIn    map[string][]int              `json:"nLabelIn,omitempty"`    // 一个 key 中的 value 值只要有一个满足就可以,数值型标签
	LabelRange  map[string]ValueRange[string] `json:"labelRange,omitempty"`  // 字符串值的范围
	NLabelRange map[string]ValueRange[int]    `json:"nLabelRange,omitempty"` // 字符串值的范围,数值型标签
	UpdateTime  ValueRange[string]            `json:"updateTime,omitempty"`  // 指定更新时间范围，格式为"yyyy-MM-dd HH:mm:ss"
	CreateTime  ValueRange[string]            `json:"createTime,omitempty"`  // 指定创建时间范围，格式为"yyyy-MM-dd HH:mm:ss"
	State       int                           `json:"state,omitempty"`       // 状态。1正常（缺省）， -1为删除
	Desc        bool                          `json:"desc,omitempty"`        // 是否倒序
	FromId      int                           `json:"fromId,omitempty"`      // 用于分页，形式如： where id > fromId
	Limit       int                           `json:"limit"`                 // 用于限制返回的条数
	Skip        int                           `json:"skip,omitempty"`        // 用于跳过指定的条数
}

type ValueRange[T comparable] struct {
	ValueGT T `json:"valueGT,omitempty"` // 值大于
	ValueGE T `json:"valueGE,omitempty"` // 值大于等于
	ValueLT T `json:"valueLT,omitempty"` // 值小于
	ValueLE T `json:"valueLE,omitempty"` // 值小于等于
}
